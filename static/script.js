const wrapper = document.querySelector('.wrapper');
const signUpLink = document.querySelector('.checkbox');

signUpLink.addEventListener('change', () => {
    if (signUpLink.checked) {
        wrapper.classList.add('animate-signIn');
        wrapper.classList.remove('animate-signUp');
    } else {
        wrapper.classList.add('animate-signUp');
        wrapper.classList.remove('animate-signIn');
    }
});
document.addEventListener('DOMContentLoaded', function () {
    document.getElementById('signupForm').addEventListener('submit', function (event) {
        event.preventDefault(); 
        var name = document.getElementsByName('username')[0].value;
        var email = document.getElementsByName('email')[0].value;
        var pswd = document.getElementsByName('password')[0].value;
        console.log('Name: ' + name);
        console.log('Email: ' + email);
        console.log('Password: ' + pswd);
        var value = [name,email,pswd].join(','); 
        document.cookie = 'keys='+value;
        document.getElementById("signupForm").reset();
    });
});
// document.addEventListener('DOMContentLoaded',function(){
//     document.getElementById('signin').addEventListener('submit',function (event){
//         event.preventDefault();
//         var uname =document.getElementsByName('uname')[0].value;
//         var password = document.getElementsByName('pswd')[0].value;
//         console.log('userName:'+ uname);
//         console.log('password:'+password);
//         var value=[uname,password].join(',');
//         document.cookie='keys='+value;
//         document.getElementById("signin").reset();
//     })
// })









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

    setInterval(nextSlide, 3000); // Change slide every 3 seconds (adjust as needed)
  
  
  });

  