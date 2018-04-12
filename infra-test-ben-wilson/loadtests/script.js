import http from "k6/http";
import { check } from "k6";

export let options = {
  vus: 10,
  duration: "5s"
};

export default function() {
  let req_coinbase = http.get("http://infra-test-ben-wilson:8081/eth_coinbase");
  check(req_coinbase, {
    "status was 200": (r) => r.status == 200,
    "transaction time OK": (r) => r.timings.duration < 100
  });

  let req_blockNumber = http.get("http://infra-test-ben-wilson:8081/eth_blockNumber");
  check(req_blockNumber, {
    "status was 200": (r) => r.status == 200,
    "transaction time OK": (r) => r.timings.duration < 100
  });
  
  let req_eth_gasPrice = http.get("http://infra-test-ben-wilson:8081/eth_gasPrice");
  check(req_eth_gasPrice, {
    "status was 200": (r) => r.status == 200,
    "transaction time OK": (r) => r.timings.duration < 100
  });
};
