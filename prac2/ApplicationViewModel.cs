using System;
using System.Collections.Generic;
using System.Collections.ObjectModel;
using System.ComponentModel;
using System.Linq;
using System.Runtime.CompilerServices;
using System.Text;
using System.Threading.Tasks;
using System.Windows;

namespace MVVM
{
    public class ApplicationViewModel : INotifyPropertyChanged
    {
        IFileService _fileService;
        IDialogService _dialogService;
        IEditorService _editorService;
        Note _selectedNote;
        public ObservableCollection<Note> Notes { get; set; }

        // команда сохранения файла
        private RelayCommand saveCommand;
        public RelayCommand SaveCommand
        {
            get
            {
                return saveCommand ??
                  (saveCommand = new RelayCommand(obj =>
                  {
                      try
                      {
                          if (_dialogService.SaveFileDialog() == true)
                          {
                              _fileService.Save(_dialogService.FilePath, Notes.ToList());
                              _dialogService.ShowMessage("Файл сохранен");
                          }
                      }
                      catch (Exception ex)
                      {
                          _dialogService.ShowMessage(ex.Message);
                      }
                  }));
            }
        }

        // команда открытия файла
        private RelayCommand openCommand;
        public RelayCommand OpenCommand
        {
            get
            {
                return openCommand ??
                  (openCommand = new RelayCommand(obj =>
                  {
                      try
                      {
                          if (_dialogService.OpenFileDialog() == true)
                          {
                              var notes = _fileService.Open(_dialogService.FilePath);
                              Notes.Clear();
                              foreach (var p in notes)
                                  Notes.Add(p);
                              _dialogService.ShowMessage("Файл открыт");
                          }
                      }
                      catch (Exception ex)
                      {
                          _dialogService.ShowMessage(ex.Message);
                      }
                  }));
            }
        } 

        // команда добавления нового объекта
        private RelayCommand addCommand;
        public RelayCommand AddCommand
        {
            get
            {
                return addCommand ??
                  (addCommand = new RelayCommand(obj =>
                  {
                      Note note = new Note("Default Title", "Default Description");
                      Notes.Insert(0, note);
                      SelectedNote = note;
                  }));
            }
        }

        private RelayCommand removeCommand;
        public RelayCommand RemoveCommand
        {
            get
            {
                return removeCommand ??
                    (removeCommand = new RelayCommand(obj =>
                    {
                        Note note = obj as Note;
                        if (note != null)
                        {
                            Notes.Remove(note);
                        }
                    },
                    (obj) => Notes.Count > 0));
            }
        }

        private RelayCommand _openEditorWindow;
        public RelayCommand OpenEditorWindow
        {
            get
            {
                return _openEditorWindow ??
                    (_openEditorWindow = new RelayCommand(obj =>
                    {
                        ShowEditWindow();
                    },
                    (obj) => _selectedNote != null));
            }
        }

        private void ShowEditWindow()
        {
            System.Windows.Window editor;
            bool error;
            (editor, error) = _editorService.EditorDialog(_selectedNote);

            if (error) { /*какой-то лог*/ }
            editor.Show();
        }


        public Note SelectedNote
        {
            get { return _selectedNote; }
            set
            {
                _selectedNote = value;
                OnPropertyChanged("SelectedNote");
            }
        }

        public ApplicationViewModel(IDialogService dialogService, IFileService fileService, IEditorService editorService)
        {
            _dialogService = dialogService;
            _fileService = fileService;
            _editorService = editorService;

            Notes = new ObservableCollection<Note>();
        }

        public event PropertyChangedEventHandler PropertyChanged;
        public void OnPropertyChanged([CallerMemberName] string prop = "")
        {
            if (PropertyChanged != null)
                PropertyChanged(this, new PropertyChangedEventArgs(prop));
        }
    }
}
