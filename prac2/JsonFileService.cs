using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Runtime.Serialization.Json;
using System.Text;
using System.Threading.Tasks;

namespace MVVM
{
    public class JsonFileService : IFileService
    {
        public List<Note> Open(string filename)
        {
            List<Note> notes = new List<Note>();
            DataContractJsonSerializer jsonFormatter =
                new DataContractJsonSerializer(typeof(List<Note>));
            using (FileStream fs = new FileStream(filename, FileMode.OpenOrCreate))
            {
                notes = jsonFormatter.ReadObject(fs) as List<Note>;
            }

            return notes;
        }

        public void Save(string filename, List<Note> notesList)
        {
            DataContractJsonSerializer jsonFormatter =
                new DataContractJsonSerializer(typeof(List<Note>));
            using (FileStream fs = new FileStream(filename, FileMode.Create))
            {
                jsonFormatter.WriteObject(fs, notesList);
            }
        }
    }
}
