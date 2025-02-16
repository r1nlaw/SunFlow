const map = L.map('map').setView([0, 0], 2); // Изначальный центр карты (0,0) и масштаб 2

// Добавление карты OpenStreetMap
L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
    attribution: '&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
}).addTo(map);

// Функция для получения данных о солнечном свете
function getSunlightData(lat, lng) {
    $.get(`/sunlight?lat=${lat}&lng=${lng}`, function(data) {
        console.log(data);
        // Отображаем данные на карте
        const sunrise = new Date(data.sunrise);
        const sunset = new Date(data.sunset);

        const popupContent = `
            <strong>Sunlight Info:</strong><br>
            Sunrise: ${sunrise.toLocaleTimeString()}<br>
            Sunset: ${sunset.toLocaleTimeString()}<br>
            Day Length: ${data.day_length}
        `;

        L.marker([lat, lng]).addTo(map)
            .bindPopup(popupContent)
            .openPopup();
    }).fail(function() {
        alert("Failed to retrieve sunlight data.");
    });
}

// Обработчик клика на карту
map.on('click', function(e) {
    const lat = e.latlng.lat;
    const lng = e.latlng.lng;
    getSunlightData(lat, lng); // Запросим данные о солнечном свете по координатам
});
