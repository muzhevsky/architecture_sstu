using Microsoft.AspNetCore.Http;
using Microsoft.AspNetCore.Mvc;

namespace BrokerTest.Controllers
{
    [Route("api/[controller]")]
    [ApiController]
    public class RabbitMqController : ControllerBase
    {

        private readonly Services.IRabbitMqService _mqService;

        public RabbitMqController(Services.IRabbitMqService mqService)
        {
            _mqService = mqService;
        }

        [Route("[action]/{message}")]
        [HttpGet]
        public IActionResult SendMessage(string message)
        {
            _mqService.SendMessage(message);
            return Ok("Сообщение отправлено");
        }
    }

}
