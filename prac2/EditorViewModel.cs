using MVVM;
using System.ComponentModel;
using System.Runtime.CompilerServices;
using System.Windows;

namespace MVVM
{
    public class EditorViewModel : INotifyPropertyChanged
    {
        private Note _note;
        private System.Windows.Window _editorScreen;

        public Note Note
        {
            get { return _note; }
            set { _note = value; }
        }
        public EditorViewModel()
        {

        }
        public EditorViewModel(Note p, Window editor)
        {
            _note = p;
            _editorScreen = editor;
        }

        public string Title
        {
            get { return _note.Title; }
            set
            {
                _note.Title = value;
                OnPropertyChanged("Title");
            }
        }
        public string Description
        {
            get { return _note.Description; }
            set
            {
                _note.Description = value;
                OnPropertyChanged("Description");
            }
        }

        private RelayCommand _closeEditorCommand;
        public RelayCommand CloseEditorCommand
        {
            get
            {
                return _closeEditorCommand ?? new RelayCommand(obj =>
            {
                _editorScreen.Close();
            });
            }
        }

        public event PropertyChangedEventHandler PropertyChanged;
        public void OnPropertyChanged([CallerMemberName] string prop = "")
        {
            if (PropertyChanged != null)
                PropertyChanged(this, new PropertyChangedEventArgs(prop));
        }
    }
}