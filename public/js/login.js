let cnt = 0;
function toggleit(){
    const signup = document.getElementById("sign_up");
    const signin = document.getElementById("sign_in");
    if ((++cnt) & 1){
        signup.style.display = 'none';
        signin.style.display = 'flex';
    }else{
        signin.style.display = 'none';
        signup.style.display = 'flex';
    }
}