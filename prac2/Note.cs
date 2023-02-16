using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Linq;
using System.Runtime.CompilerServices;
using System.Text;
using System.Threading.Tasks;

namespace MVVM
{
    public class Note
    {
        private string _title;
        private string _description;

        public string Title
        {
            get { return _title; }
            set { _title = value; }
        }
        public string Description
        {
            get { return _description; }
            set { _description = value; }
        }

        public Note(string title, string description)
        {
            _title = title; 
            _description = description;
        }
    }
}
