using Microsoft.AspNetCore.Http;
using Microsoft.AspNetCore.Mvc;
using Newtonsoft.Json;
using System.Text.Json.Serialization;

namespace prac4.Controllers
{
    [Route("api/[controller]")]
    [ApiController]
    public class PickDeliveryController : ControllerBase
    {
        [HttpPut(Name = "PickDelivery")]
        public string PickDelivery([FromQuery(Name = "PickDelivery")] int deliveryId)
        {
            return $"Заказ под номером {deliveryId} принят";
        }

    }
}
