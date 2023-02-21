using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Windows;

namespace MVVM
{
    public class DefaultEditorService : IEditorService
    {
        public (Window, bool) EditorDialog(Note note)
        {
            return (new Window(note), false);
        }
    }
}
