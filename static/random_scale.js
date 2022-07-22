const app = Vue.createApp({})
    app.component('scale-demonstration', {
        data() {
            return {
                currentScale: {
                    ScaleName: '',
                    ModeName: '', 
                    Notes: [],
                },
                keys: [
                    {note: 'C_1', sharp: false},
                    {note: 'Db_1', sharp: true},
                    {note: 'D_1', sharp: false},
                    {note: 'Eb_1', sharp: true},
                    {note: 'E_1', sharp: false},
                    {note: 'F_1', sharp: false},
                    {note: 'Gb_1', sharp: true},
                    {note: 'G_1', sharp: false},
                    {note: 'Ab_1', sharp: true},
                    {note: 'A_1', sharp: false},
                    {note: 'Bb_1', sharp: true},
                    {note: 'B_1', sharp: false},
                    {note: 'C_2', sharp: false},
                    {note: 'Db_2', sharp: true},
                    {note: 'D_2', sharp: false},
                    {note: 'Eb_2', sharp: true},
                    {note: 'E_2', sharp: false},
                    {note: 'F_2', sharp: false},
                    {note: 'Gb_2', sharp: true},
                    {note: 'G_2', sharp: false},
                    {note: 'Ab_2', sharp: true},
                    {note: 'A_2', sharp: false},
                    {note: 'Bb_2', sharp: true},
                    {note: 'B_2', sharp: false},
                ],
                keyValues: {'C': 0, 'Db': 1, 'D': 2, 'Eb': 3, 'E': 4, 'F': 5, 'Gb': 6, 'G': 7, 'Ab': 8, 'A': 9, 'Bb' : 10, 'B': 11}
            }
        },
        methods:{
            getScale() {
                fetch('/random_scale')
                    .then(response => response.json())
                    .then(data => {
                            this.currentScale.ScaleName = data.ScaleName
                            this.currentScale.ModeName = data.ModeName
                            this.currentScale.Notes = data.Notes
                        })
            },
            getNoteInfo(note) {
                let noteName
                let octaveNumber = 1
                const firstNoteValue = this.keyValues[this.currentScale.Notes[0]]
                for(const index of this.currentScale.Notes.keys()){
                    noteName = this.currentScale.Notes[index]
                    
                    if(octaveNumber === 1 && index !== 0 && this.keyValues[noteName] <= firstNoteValue){
                        octaveNumber += 1
                    }
                    numberedNoteName = noteName + "_" + octaveNumber
                    if(note === numberedNoteName){
                        return {
                            inScale: true,
                            index: index
                        }
                    }
                }
                return {
                    inScale: false,
                    index: -1
                }
            },
            getBackgroundColor(noteIndex){
                if(noteIndex === -1){
                    return
                }
                const colors = {
                    0 : '#eb4034',
                    1 : '#eb5c34',
                    2 : '#eb6534',
                    3 : '#eb8c34',
                    4 : '#eb9934',
                    5 : '#eba534',
                    6 : '#ebb434',
                    7 : '#ebc334',
                    8 : '#ebd934',
                    9 : '#ebeb34',
                    10: '#cdeb34',
                    11: '#b7eb34',
                    12: '#9feb34',
                }
                return {'background' : colors[noteIndex]}
            }
        },
        template: `
            <div class="keys">
                <div v-for="key in keys"
                    class="key" 
                    :class="{
                        'sharp': key.sharp,
                        'highlighted': this.getNoteInfo(key.note).inScale
                    }"
                    v-bind:style="this.getBackgroundColor(this.getNoteInfo(key.note).index)"
                    v-bind:data-note="key.note"
                    v-bind:data-index="this.getNoteInfo(key.note).index"
                >
                    <div v-if="this.getNoteInfo(key.note).index !== -1" class="noteIndex">{{ this.getNoteInfo(key.note).index+1 }}</div>
                </div>
            </div>
            <div class="get-random-scale">
                <div class="scale" v-if="this.currentScale.ScaleName !== ''">
                    <div class="scale-title">Scale</div>
                    <div class="scale-name">{{this.currentScale.ScaleName}}</div>
                </div>
                <button class="get-scale" v-on:click="getScale">Get Random Scale</button>
                <div class="mode" v-if="this.currentScale.ModeName !== ''">
                    <div class="mode-title">Mode</div>
                    <div class="mode-name">{{this.currentScale.ModeName}}</div>
                </div>
            </div>
        `
     })
     app.mount('#app')