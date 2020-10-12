const url = "127.0.0.1:8000"

var inputForm = document.getElementById("inputForm")

inputForm.addEventListener("submit", (e) => {
    e.preventDefault()

    const formdata = new FormData(inputForm)
    fetch(url, {
        method:"POST",
        body:formdata,
    })
})