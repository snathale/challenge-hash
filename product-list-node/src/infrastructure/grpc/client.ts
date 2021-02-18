import { credentials } from "grpc"
import config from "../../app/config";
import { CalculatorClient } from "./calculator_grpc_pb";

export default new CalculatorClient(config.calculatorServiceHost,credentials.createInsecure());