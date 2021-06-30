import { RepoItem } from './repo-item';
import { Tag } from './tag';

export interface OwnerTrimmed {
    id: number;
    name: string;
    avatar: string;
}

export interface TimeEntry {
    id?: number;
    user: OwnerTrimmed;
    organisation?: OwnerTrimmed;
    comments: string;
    created?: string;
    updated?: string;
    value: number;
    valueType: string;
    tags?: Tag[];
    repoItems?: RepoItem[];
}
