using System;

namespace prac4
{
    public class Courier : IEquatable<Courier>
    {
        public LocationData Location { get; private set; }
        public int Id { get; private set; }
        private static List<Courier> Couriers = new List<Courier>();

        public static Courier GetCourier(int id)
        {
            foreach(var item in Couriers)
                if (item.Id == id) return item;

            return null;
        }

        public Courier(int id) {
            Id = id;
            Location = new LocationData();
            Location.X = 0.5f;
            Location.Y = 5.2f;
            Couriers.Add(this);
        }

        public bool Equals(Courier? other)
        {
            if (other == null) return false;
            return this.Id == other.Id;
        }
    }
}
