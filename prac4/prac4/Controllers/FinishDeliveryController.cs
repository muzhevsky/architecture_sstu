using Microsoft.AspNetCore.Http;
using Microsoft.AspNetCore.Mvc;

namespace prac4.Controllers
{
    [Route("api/[controller]")]
    [ApiController]
    public class FinishDeliveryController : ControllerBase
    {
        [HttpPut(Name = "FinishDelivery")]
        public string FinishDelivery([FromQuery(Name = "FinishDelivery")] int deliveryId)
        {
            return $"доставка {deliveryId} успешна";
        }
    }
}
