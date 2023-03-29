export interface BeLiveInfo {
    id : number
    username :string
	photo :string
	img :string
	title :string
	online : number
}

export type GetBeLiveListRes = Array<BeLiveInfo>