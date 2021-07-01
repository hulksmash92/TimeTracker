import { Component, OnInit } from '@angular/core';
import { Tag } from 'src/app/models/tag';

import { TimeService } from 'src/app/services/time/time.service';

@Component({
  selector: 'time-form',
  templateUrl: './form.component.html',
  styleUrls: ['./form.component.scss']
})
export class FormComponent implements OnInit {
  tags: Tag[] = [];
  readonly valueTypes: string[] = ['Hours', 'Minutes', 'Units'];

  constructor(private readonly timeService: TimeService) { }

  ngOnInit(): void {
    this.setTags();
  }

  setTags(): void {
    this.timeService.getTags()
      .subscribe((res: Tag[]) => {
        this.tags = res || [];
      });
  }

  submit(): void {

  }

}
