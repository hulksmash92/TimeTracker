import { Component, OnInit, Output, EventEmitter } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { RepoItem } from 'src/app/models/repo-item';

import { Tag } from 'src/app/models/tag';
import { TimeService } from 'src/app/services/time/time.service';

@Component({
  selector: 'time-form',
  templateUrl: './form.component.html',
  styleUrls: ['./form.component.scss']
})
export class FormComponent implements OnInit {
  @Output() cancel: EventEmitter<void> = new EventEmitter<void>();
  readonly valueTypes: string[] = ['Hours', 'Minutes', 'Units'];
  tagOptions: Tag[] = [];
  repo: any;
  repoItems: RepoItem[] = [];
  selectedTags: Tag[] = [];

  formGroup: FormGroup = new FormGroup({
    comments: new FormControl(null, [Validators.required]),
    value: new FormControl(null, [Validators.required]),
    valueType: new FormControl('Hours', [Validators.required]),
  });

  constructor(private readonly timeService: TimeService) { }

  ngOnInit(): void {
    this.setTagOptions();
  }

  setTagOptions(): void {
    this.timeService.getTags()
      .subscribe((res: Tag[]) => {
        this.tagOptions = res || [];
      });
  }

  setRepoItems(items: RepoItem[]): void {
    items = items.filter(i => {
      const index = this.repoItems.findIndex(r => r.itemIdSource === i.itemIdSource && r.itemType === i.itemType);
      return index === -1;
    });
    this.repoItems.push(...items);
  }

  submit(): void {
    if (this.formGroup.valid) {

    }
  }

  resetForm(): void {
    this.formGroup.reset({ 
      comments: null, 
      value: null,
      valueType: 'Hours'
    });
  }

}
