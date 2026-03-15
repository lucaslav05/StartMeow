const { loadEnvFile } = require('node:process');
loadEnvFile();

db = connect(process.env.MONGODB_URI);

console.log(db);