import http from 'k6/http';
import { check, sleep } from 'k6';

export let options = {
    vus: 100, // Number of virtual users
    duration: '30s', // Duration of the test
};

const BASE_URL = 'http://localhost:8087';

export default function () {
    group("healthz", () => {
        let res = http.get(`${BASE_URL}/healthz/liveness`);
        check(res, {
            'liveness is status 200': (r) => r.status === 200,
        });

        res = http.get(`${BASE_URL}/healthz/readiness`);
        check(res, {
            'readiness is status 200': (r) => r.status === 200,
        });
    });

    res = http.get(`${BASE_URL}/`);
    check(res, {
        'main page loads': (r) => r.status === 200,
    });

    let payload = { item_id: '1', quantity: '2' };
    res = http.post(`${BASE_URL}/add-item`, payload);
    check(res, {
        'add item redirects': (r) => r.status === 302,
    });

    res = http.get(`${BASE_URL}/remove-cart-item?cart_item_id=1`);
    check(res, {
        'remove item redirects': (r) => r.status === 302,
    });

    sleep(1);
}
