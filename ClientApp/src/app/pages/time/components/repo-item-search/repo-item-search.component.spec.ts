import { ComponentFixture, TestBed } from '@angular/core/testing';

import { RepoItemSearchComponent } from './repo-item-search.component';

describe('RepoItemSearchComponent', () => {
  let component: RepoItemSearchComponent;
  let fixture: ComponentFixture<RepoItemSearchComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ RepoItemSearchComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(RepoItemSearchComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
