import { Component } from '@angular/core';

@Component({
  selector: 'app-settings',
  templateUrl: './settings.component.html',
  styleUrls: ['./settings.component.scss']
})
export class SettingsComponent {
  /**
   * Name of the section currently in the view
   */
  visibleSection: string = 'profile';

  /**
   * Changes the visible section in the view
   * @param section name of the section to view
   */
  changeVisibleSection(section: string): void {
    if (this.visibleSection !== section) {
      this.visibleSection = section;
    }
  }
}
