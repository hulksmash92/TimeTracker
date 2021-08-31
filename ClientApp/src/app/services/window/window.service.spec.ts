import { TestBed } from '@angular/core/testing';

import { WindowService } from './window.service';

describe('WindowService', () => {
  let service: WindowService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(WindowService);
    service.window = {
      location: {
        href: 'http://localhost:5000/'
      }
    } as any;
  });

  describe('#goExternal()', () => {
    const getHref = () => service.window.location.href;

    it('should set #window.location.href to url param value when truthy and starts with http', () => {
      service.goExternal('https://www.example.com');
      expect(getHref()).toEqual('https://www.example.com');

      service.goExternal('http://www.example.com');
      expect(getHref()).toEqual('http://www.example.com');
    });

    it('should not set #window.location.href to url param value when falsy or doesn\'t starts with http', () => {
      const initVal = 'http://localhost:5000/';

      service.goExternal('www.example.com');
      expect(getHref()).toEqual(initVal);

      service.goExternal('');
      expect(getHref()).toEqual(initVal);

      service.goExternal(null);
      expect(getHref()).toEqual(initVal);

      service.goExternal(undefined);
      expect(getHref()).toEqual(initVal);
    });

  });
});
