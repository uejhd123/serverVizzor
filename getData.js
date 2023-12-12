setInterval(async function() {
    // Получаем данные от сервера
    data = await fetch("/getData").then(response => response.json());
    console.log(data)
    // Выводим данные в HTML
    document.getElementById("data").innerHTML = data.cpu;
}, 1000);
