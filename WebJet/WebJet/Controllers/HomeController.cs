using System;
using System.Collections.Generic;
using System.Diagnostics;
using System.Linq;
using System.Net.Http;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;
using Newtonsoft.Json;
using WebJet.Models;

namespace WebJet.Controllers
{
    public class HomeController : Controller
    {
        private static readonly HttpClient Client = new HttpClient();
        public async Task<IActionResult> Index()
        {
            List<MovieVm> account = new List<MovieVm>();
            HttpResponseMessage response =  await Client.GetAsync("http://localhost:8081//api/movies/");
            if (response.IsSuccessStatusCode)
            {
               string product = await response.Content.ReadAsStringAsync();
                account = JsonConvert.DeserializeObject<List<MovieVm>>(product);
            }
            return View(account);
        }

        public async Task<IActionResult> Detail(int movieId,string provider)
        {
            MoviePVm account = new MoviePVm();
            
                HttpResponseMessage response =  await Client.GetAsync("http://localhost:8081//api/"+provider+"/movie/"+movieId.ToString());
                if (response.IsSuccessStatusCode)
                {
                    string product = await response.Content.ReadAsStringAsync();
                    account = JsonConvert.DeserializeObject<MoviePVm>(product);
                }
                return View(account);
        }

        public async Task<IActionResult> Cinema()
        {
            List<MovieVm> account = new List<MovieVm>();
            HttpResponseMessage response =  await Client.GetAsync("http://localhost:8081/api/cinemaworld/movies");
            if (response.IsSuccessStatusCode)
            {
                string product = await response.Content.ReadAsStringAsync();
                account = JsonConvert.DeserializeObject<List<MovieVm>>(product);
            }
            return View(account);
        }
        
        public async Task<IActionResult> Film()
        {
            List<MovieVm> account = new List<MovieVm>();
            HttpResponseMessage response =  await Client.GetAsync("http://localhost:8081/api/filmworld/movies");
            if (response.IsSuccessStatusCode)
            {
                string product = await response.Content.ReadAsStringAsync();
                account = JsonConvert.DeserializeObject<List<MovieVm>>(product);
            }
            return View(account);
        }
        
        public IActionResult Contact()
        {
            ViewData["Message"] = "Your contact page.";

            return View();
        }

        public IActionResult Error()
        {
            return View(new ErrorViewModel {RequestId = Activity.Current?.Id ?? HttpContext.TraceIdentifier});
        }
    }
}