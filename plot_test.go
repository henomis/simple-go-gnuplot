package plot

import "testing"

func TestSimpleGnuplot_buildGnuplotCommand(t *testing.T) {
	type fields struct {
		gnuplotExecutable    string
		environmentVariables map[string]interface{}
		plotFilePath         string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "test #1",
			fields: fields{
				gnuplotExecutable:    "/bin/gnuplot",
				environmentVariables: make(map[string]interface{}),
				plotFilePath:         "/tmp/plot01.plot",
			},
			want: "/bin/gnuplot /tmp/plot01.plot",
		},
		{
			name: "test #2",
			fields: fields{
				gnuplotExecutable: "/bin/gnuplot",
				environmentVariables: map[string]interface{}{
					"key1": "value1",
					"key2": 2,
				},
				plotFilePath: "/tmp/plot01.plot",
			},
			want: "/bin/gnuplot -e \"key1='value1'; key2='2'\" /tmp/plot01.plot",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sg := &simpleGnuplot{
				gnuplotExecutablePath: tt.fields.gnuplotExecutable,
				environmentVariables:  tt.fields.environmentVariables,
				plotFilePath:          tt.fields.plotFilePath,
			}
			if got := sg.buildGnuplotCommand(); got != tt.want {
				t.Errorf("SimpleGnuplot.buildGnuplotCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}
