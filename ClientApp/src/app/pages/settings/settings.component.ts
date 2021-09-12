import { Component } from '@angular/core';

@Component({
  selector: 'app-settings',
  templateUrl: './settings.component.html',
  styleUrls: ['./settings.component.scss']
})
export class SettingsComponent {
  visibleSection: string = 'profile';

  
  changeVisibleSection(section: string): void {
    if (this.visibleSection !== section) {
      this.visibleSection = section;
    }
  }
}
