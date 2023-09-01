import http from 'k6/http';
import { sleep } from 'k6';

export const options = {
    vus: 3, 
    duration: '1m',
  };

export default function () {
  http.get('https://test.k6.io', {
    tags: { name: 'TrackEvent' },
  });
  sleep(1);
}