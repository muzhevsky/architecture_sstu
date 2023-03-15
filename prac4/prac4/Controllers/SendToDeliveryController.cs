using Microsoft.AspNetCore.Http;
using Microsoft.AspNetCore.Mvc;

namespace prac4.Controllers
{
    [Route("api/[controller]")]
    [ApiController]
    public class SendToDeliveryController : ControllerBase
    {
        [HttpPut(Name = "SendToDelivery")]
        public string SendToDelivery([FromQuery(Name = "SendToDelivery")] int deliveryId)
        {
            return $"доставка {deliveryId} отправлена службе доставки";
        }
    }
}
