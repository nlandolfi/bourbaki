package bbk

const MakefileTemplate = `# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make {{ .Config.Name }}.pdf
{{ if .HasLitFile }}
sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex
{{ end }}
{{ .Config.Name }}.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > {{ .Config.Name }}.tex

{{ .Config.Name }}.pdf: ../../*.tex ../../trademark.pdf *.tex {{ .Config.Name }}.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode {{ .Config.Name }}.tex
	../../latexrun {{ .Config.Name }}.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: {{ .Config.Name }}.pdf
	open {{ .Config.Name }}.pdf

o:
	make open

terms: {{ .Config.Name }}.tex
	bbk_sheet -mode ts {{ .Config.Name }}.tex

spell:
	aspell -c sheet.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f {{.Config.Name}}.*
	make reset
	bbk_sheet -mode mk > Makefile`

// TODO: eventually change the generated header, but leave for now so git diffs are small
const MakefileTemplate2 = `# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make {{ .Config.Name }}.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

{{ .Config.Name }}.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > {{ .Config.Name }}.tex

{{ .Config.Name }}.pdf: ../../*.tex ../../trademark.pdf *.tex {{ .Config.Name }}.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode {{ .Config.Name }}.tex
	../../latexrun {{ .Config.Name }}.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: {{ .Config.Name }}.pdf
	open {{ .Config.Name }}.pdf

o:
	make open

terms: {{ .Config.Name }}.tex
	bbk_sheet -mode ts {{ .Config.Name }}.tex

spell:
	aspell -c sheet.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f {{.Config.Name}}.*
	make reset
	bbk_sheet -mode mk > Makefile`
