import cors from 'cors';
import * as dotenv from 'dotenv';
import App from './app/app';

dotenv.config();
const { JSend } = require('jsend-express')
const jSend = new JSend({ name: 'appName', version: 'X.X.X', release: 'XX' })
const app = new App()
app.server.use(jSend.middleware.bind(jSend))
app.server.use(cors())
app.run()