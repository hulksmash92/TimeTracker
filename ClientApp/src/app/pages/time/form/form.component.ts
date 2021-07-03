import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';

import { Tag } from 'src/app/models/tag';
import { RepoService } from 'src/app/services/repo/repo.service';
import { TimeService } from 'src/app/services/time/time.service';

@Component({
  selector: 'time-form',
  templateUrl: './form.component.html',
  styleUrls: ['./form.component.scss']
})
export class FormComponent implements OnInit {
  readonly valueTypes: string[] = ['Hours', 'Minutes', 'Units'];
  tags: Tag[] = [];
  repo: any;
  

  constructor(private readonly timeService: TimeService, private readonly repoService: RepoService) { }

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
