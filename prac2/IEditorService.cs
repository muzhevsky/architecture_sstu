using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Windows;

namespace MVVM
{
    public interface IEditorService
    {
        (Window, bool) EditorDialog(Note note);
    }
}
