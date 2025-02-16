// Функция для обновления карты на основе солнечного света
function showSunlight() {
    $.getJSON("/sunlight", function(data) {
      // Преобразуем время восхода и заката в объекты Date
      var sunrise = new Date('1970-01-01T' + data.sunrise + 'Z');
      var sunset = new Date('1970-01-01T' + data.sunset + 'Z');
      
      // Получаем текущее время
      var now = new Date();
  
      // Логика для расчета светлой и темной зоны
      var isDaytime = now > sunrise && now < sunset;
      
      // Если день, отображаем светлую часть, если ночь - темную
      if (isDaytime) {
        L.polygon([
          [90, -180], [90, 0], [0, 0], [0, -180]
        ], { color: 'yellow' }).addTo(map); // Светлая зона
      } else {
        L.polygon([
          [90, 0], [90, 180], [0, 180], [0, 0]
        ], { color: 'gray' }).addTo(map); // Темная зона
      }
    });
  }
  
  showSunlight(); // Показываем солнечное освещение
  