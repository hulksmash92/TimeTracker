import { Component, EventEmitter, Input, Output } from '@angular/core';
import { FormControl } from '@angular/forms';
import { COMMA, ENTER } from '@angular/cdk/keycodes';
import { MatChipInputEvent } from '@angular/material/chips';

import { Observable } from 'rxjs';
import { map } from 'rxjs/operators';

import { Tag } from 'src/app/models/tag';

@Component({
  selector: 'time-form-tags',
  templateUrl: './time-form-tags.component.html',
  styleUrls: ['./time-form-tags.component.scss']
})
export class TimeFormTagsComponent {
  @Input() autocompleteTags: Tag[] = [];
  @Input() selectedTags: Tag[] = [];
  @Output() tagAdded: EventEmitter<Tag> = new EventEmitter<Tag>();
  @Output() selectedTagsChanged: EventEmitter<Tag[]> = new EventEmitter<Tag[]>();
  tagFc: FormControl = new FormControl();
  separatorKeysCodes: number[] = [ENTER, COMMA];
  filteredTags: Observable<Tag[]>;

  constructor() {
    this.filteredTags = this.tagFc.valueChanges.pipe(
      map((v: string) => !!v ? this.filterTags(v) : this.autocompleteTags.slice())
    );
  }

  filterTags(value: string): Tag[] {
    const filterVal = (value || '').toLowerCase();
    return this.autocompleteTags
      .filter(t => t.name.toLowerCase().includes(filterVal));
  }

  /**
   * Creates a tag object and emits it to the parent to add to the list
   */
  addTag(event: MatChipInputEvent): void {
    const tagName: string = (event.value || '').trim();

    if (!!tagName) {
      const selected = this.selectedTags.findIndex(t => t.name.toLowerCase() === tagName.toLowerCase()) > -1;

      if (!selected) {
        let tag: Tag = this.autocompleteTags.find(t => t.name.toLowerCase() === tagName.toLowerCase());
        if (!tag) {
          tag = { name: tagName };
        }
        this.tagAdded.emit(tag);
      }
    }
    
    event.chipInput!.clear();
    this.tagFc.setValue(null);
  }

  /**
   * Removes a tag from the selected tag list and then emits the new list to the parent
   */
  removeTag(tag: Tag): void {
    const index = this.selectedTags.findIndex(t => t.name === tag.name);
    if (index > -1) {
      this.selectedTags.splice(index, 1);
      this.selectedTagsChanged.emit(this.selectedTags);
    }
  }

}
