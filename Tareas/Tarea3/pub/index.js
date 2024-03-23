const Redis = require('ioredis');
const conexion = new Redis({
host: '10.115.137.67',
port: 6379,
connectTimeout: 5000,
});

function funcionPub() {
	conexion.publish('test', 'hola a todos')
	.then(() => {
		console.log('Mensaje publicado');
	})
	.catch(err => {
		console.log('eror al publicar: ',err);
	});

}

setInterval(funcionPub, 3000);
