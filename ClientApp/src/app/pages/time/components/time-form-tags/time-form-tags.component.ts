import { Component, EventEmitter, Input, Output, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';
import { COMMA, ENTER } from '@angular/cdk/keycodes';
import { MatChipInputEvent } from '@angular/material/chips';

import { Tag } from 'src/app/models/tag';
import { MatAutocompleteSelectedEvent } from '@angular/material/autocomplete';

@Component({
  selector: 'time-form-tags',
  templateUrl: './time-form-tags.component.html',
  styleUrls: ['./time-form-tags.component.scss']
})
export class TimeFormTagsComponent implements OnInit {
  @Input() set autocompleteTags(v: Tag[]) {
    if (!v) {
      v = [];
    }
    this.availableTags = v;
    this.filteredTags = this.filterTags('');
  }
  @Input() selectedTags: Tag[] = [];
  @Output() tagAdded: EventEmitter<Tag> = new EventEmitter<Tag>();
  @Output() selectedTagsChanged: EventEmitter<Tag[]> = new EventEmitter<Tag[]>();
  tagFc: FormControl = new FormControl();
  separatorKeysCodes: number[] = [ENTER, COMMA];
  availableTags: Tag[] = [];
  filteredTags: Tag[] = [];

  ngOnInit(): void {
    this.tagFc.valueChanges.subscribe({
      next: (v: string) => {
        this.filteredTags = this.filterTags(v);
      }
    });
  }

  filterTags(value: string): Tag[] {
    const filterVal = (value || '').toLowerCase();
    return (this.availableTags || []).filter(t => t.name.toLowerCase().includes(filterVal));
  }

  /**
   * Creates a tag object and emits it to the parent to add to the list
   */
  addTag(event: MatChipInputEvent): void {
    const tagName: string = (event.value || '').trim();
    
    if (!!tagName) {
      const selected = this.selectedTags.findIndex(t => t.name.toLowerCase() === tagName.toLowerCase()) > -1;

      if (!selected) {
        let tag = this.availableTags.find(t => t.name.toLowerCase() === tagName.toLowerCase());
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
   * Emits the selected tag to the parent component
   * @param event Event received when an autocomplete option is selected
   */
  autocompleteOptionSelect(event: MatAutocompleteSelectedEvent): void {
    const tag: Tag = event.option.value;
    this.tagAdded.emit(tag);
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
