апишка для сбора статы по дискам со стораджей
сейчас поддерживает 3 метода, POST, GET и DELETE

структура json ожидающегося на вход описана в models/diskstat.go 

юзать так:
добавит объект, если успешно то возвращает ID из коллекции в mongodb и сообщение
со статусом
```
curl -X POST -d @example.disks.json  http://127.0.0.1:3000/disks
{"ID":"5c0afe15eff0ef6d7e5589b1","Status":"OK"}
```

получить объект по его ID
```
curl -X GET  http://127.0.0.1:3000/disks/5c0b00e6eff0ef6f99443461
{
   "id":"5c0b00e6eff0ef6f99443461",
   "host":"exmpl",
    "stat":[
        {
        "mount":["/storage/20699","/storage/20700"],
        "path":"/dev/sdaaaaa","smart":...........
```

удалить объект по его ID
```
curl -v -X DELETE http://127.0.0.1:3000/remove/5c0b00e6eff0ef6f99443461
{"result":"Delete success"}
```
