import { SimpleChanges } from '@angular/core';
import { Component, Input, OnChanges } from '@angular/core';

@Component({
  selector: 'app-avatar',
  templateUrl: './avatar.component.html',
  styleUrls: ['./avatar.component.scss']
})
export class AvatarComponent implements OnChanges {
  @Input() avatar: string;
  @Input() name: string;
  @Input() diameter: string = '40px';
  initials: string;
  fontSize: string = '1.0rem';
  wrapperStyles: any = {};

  ngOnChanges(changes: SimpleChanges): void {
    const name: string = changes.name?.currentValue;
    if (!changes.avatar?.currentValue && !!name) {
      this.setInitials(name);
    }

    if (!!changes.diameter) {
      this.wrapperStyles = {
        height: this.diameter,
        width: this.diameter,
      };

      this.fontSize = this.getFontSize();
    }
  }

  /**
   * Removes and special chars from the name, then gets the first character
   * of the first two words in the name param, or first word if only one
   * word is present
   * @param name name value to clean and split
   */
  setInitials(name: string): void {
    const nameSplit = name.replace(/[^a-zA-Z0-9 ]/g, '').split(' ');
    this.initials = nameSplit[0].charAt(0);

    if (nameSplit.length > 1) {
      this.initials += nameSplit[1].charAt(0);
    }
    this.initials = this.initials.toUpperCase();
  }

  /**
   * Calculates the font-size based on the diameter using the units specified in the diameter
   */
  getFontSize(): string {
    const extracted = /(cm|mm|in|px|ex|em|ch|rem|vw|vh|vmin|vmax|%)/igm.exec(this.diameter);
    const units = !!extracted && extracted.length > 0 ? extracted[0].toLowerCase() : 'px';
    const value = parseFloat(this.diameter) * 0.55;
    return Math.floor(value) + units;
  }

}
