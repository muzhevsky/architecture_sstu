using Microsoft.AspNetCore.Http;
using Microsoft.AspNetCore.Mvc;
using Newtonsoft.Json;

namespace prac4.Controllers
{
    [Route("api/[controller]")]
    [ApiController]
    public class UpdateLocationController : ControllerBase
    {
        [HttpPut(Name = "UpdateLocationController")]
        public string UpdateLocation([FromQuery(Name = "UpdateLocationController")] int courierId)
        {
            var courier = Courier.GetCourier(courierId);
            if (courier != null)
            {
                var data = JsonConvert.SerializeObject(courier.Location);
                return $"Новая локация курьера: {data}";
            }

            return "Ошибка";
        }
    }
}
