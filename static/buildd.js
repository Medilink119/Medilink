document.addEventListener('DOMContentLoaded', function() {
    const typed = new Typed('.multiple-text', {
      strings:['<h2>Linking Lives,</h2><h4>Stitching Health, A Digital Era by</h4>','<h2>Healing Beyond Horizons,</h2><h4>Innovating and Elevating Lives by</h4>','<h2>Caring Today,</h2><h4> for a Healthier Tomorrow by</h4>','<h2>A Promise of Healing,</h2><h4>Guiding Your Journey to Wellness by</h4>'],
      typeSpeed: 100,
      backSpeed: 100,
      backDelay: 5000,
      loop: true,
      showCursor: true,
    });
  });
  var icon = document.getElementById("theme-icon");
  var body = document.body;
  var iconElement = document.getElementById("theme-icon");
  var logoIcon = document.querySelector(".logo i");
  var imgElement = document.querySelector(".slide img");
  var bg=document.querySelector(".container2-1")
  var bg1=document.querySelector(".container3-2")
  
  icon.onclick = function () {
      body.classList.toggle("dark-theme");

      if (body.classList.contains("dark-theme")) {
          iconElement.innerHTML = "<i class='bx bxs-sun' style='color:#190019'></i>";
          logoIcon.style.color = "#190019";
          imgElement.src = "media/doc2.png";
          // For bg
          var existingVideoBg = bg.querySelector("video");
          if (existingVideoBg) {
              existingVideoBg.remove();
          }
          var videoBg = document.createElement("video");
          videoBg.src = "media/white.mp4";
          videoBg.autoplay = true;
          videoBg.muted = true;
          videoBg.loop = true;
          bg.appendChild(videoBg);

          // For bg1
          var existingVideoBg1 = bg1.querySelector("video");
          if (existingVideoBg1) {
              existingVideoBg1.remove();
          }
          var videoBg1 = document.createElement("video");
          videoBg1.src = "media/white.mp4";
          videoBg1.autoplay = true;
          videoBg1.muted = true;
          videoBg1.loop = true;
          bg1.appendChild(videoBg1);

         
      } else {
          iconElement.innerHTML = "<i class='bx bxs-moon' style='color:white'></i>";
          logoIcon.style.color = "#fbe4d8";
          imgElement.src = "media/doc.png";
          var existingVideo = bg.querySelector("video");
          if (existingVideo) {
              existingVideo.remove();
          }
          var video = document.createElement("video");
          video.src = "media/black.mp4";
          video.autoplay = true;
          video.muted = true;
          video.loop = true;
          bg.appendChild(video);

          var existingVideoBg1 = bg1.querySelector("video");
          if (existingVideoBg1) {
              existingVideoBg1.remove();
          }
          var videoBg1 = document.createElement("video");
          videoBg1.src = "media/black.mp4";
          videoBg1.autoplay = true;
          videoBg1.muted = true;
          videoBg1.loop = true;
          bg1.appendChild(videoBg1);
          
      }   
  };
      document.addEventListener("DOMContentLoaded", function () {
      const slider = document.querySelector(".slider");
      let currentIndex = 0;

      function showSlide(index) {
      const translateValue = -index * 100 + "%";
      slider.style.transform = "translateX(" + translateValue + ")";
      }

      function nextSlide() {
      currentIndex = (currentIndex + 1) % slider.children.length;
      showSlide(currentIndex);
      }

      function prevSlide() {
      currentIndex = (currentIndex - 1 + slider.children.length) % slider.children.length;
      showSlide(currentIndex);
      }

      setInterval(nextSlide, 5000); 
  });

  document.addEventListener('DOMContentLoaded', function () {

    const hamburger = document.getElementById('hamburger');
        const navbar = document.getElementById('navbar');
    
        hamburger.addEventListener('click', function () {
            navbar.classList.toggle('active');
            const closeBtn = document.getElementById('closeBtn');
            if (navbar.classList.contains('active')) {
                closeBtn.style.display = 'block'; // Show the close button when the menu is open
            } else {
                closeBtn.style.display = 'none'; // Hide the close button when the menu is closed
            }
        });
    
        function closeHamburger() {
            navbar.classList.remove('active');
            const closeBtn = document.getElementById('closeBtn');
            closeBtn.style.display = 'none'; // Hide the close button
            }
  });