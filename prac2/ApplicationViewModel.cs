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
        IFileService fileService;
        IDialogService dialogService;
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
                          if (dialogService.SaveFileDialog() == true)
                          {
                              fileService.Save(dialogService.FilePath, Notes.ToList());
                              dialogService.ShowMessage("Файл сохранен");
                          }
                      }
                      catch (Exception ex)
                      {
                          dialogService.ShowMessage(ex.Message);
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
                          if (dialogService.OpenFileDialog() == true)
                          {
                              var notes = fileService.Open(dialogService.FilePath);
                              Notes.Clear();
                              foreach (var p in notes)
                                  Notes.Add(p);
                              dialogService.ShowMessage("Файл открыт");
                          }
                      }
                      catch (Exception ex)
                      {
                          dialogService.ShowMessage(ex.Message);
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
            var editor = new Editor(SelectedNote);
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

        public ApplicationViewModel(IDialogService dialogService, IFileService fileService)
        {
            this.dialogService = dialogService;
            this.fileService = fileService;

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
