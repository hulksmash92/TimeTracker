import { Injectable } from '@angular/core';

export function _window(): Window {
  return window;
}

@Injectable({
  providedIn: 'root'
})
export class WindowService {
  window: Window;
  
  constructor() {
    this.window = _window();
  }
  
  /**
   * Goes to an external url if its valid
   */
  goExternal(url: string): void {
    if (!!url && /^https?:\/\//.test(url)) {
      this.window.location.href = url;
    }
  }

}
