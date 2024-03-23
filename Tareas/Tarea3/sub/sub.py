import redis
import time

conexion = redis.StrictRedis(
    host='10.115.137.67',
    port=6379,
    socket_timeout=5
)
pubsub = conexion.pubsub()

def manejar_mensaje(message):
    print(f"Mensaje recibido en el canal '{message['channel'].decode()}' : {message['data'].decode()}")

try:
    pubsub.subscribe(**{'test': manejar_mensaje})
    canales_suscritos = pubsub.channels
    print(f"Suscrito a {len(canales_suscritos)} canal(es).")
    while True:
        try:
            for mensaje in pubsub.listen():
                pass
        except redis.exceptions.TimeoutError:
            print("Tiempo de espera alcanzado. Reintentando la conexión en 5 segundos...")
            time.sleep(5)
            pubsub.close()
            conexion = redis.StrictRedis(
                host='10.115.137.67',
                port=6379,
                socket_timeout=5
            )
            pubsub = conexion.pubsub()
            pubsub.subscribe(**{'test': manejar_mensaje})
except Exception as e:
    print("Error en el subscriptor Redis:", e)
finally:
    conexion.close()
