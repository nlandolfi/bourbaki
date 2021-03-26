package bbk

const MakefileTemplate = `# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make {{ .Name }}.pdf
	make graph.pdf

{{ .Name }}.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > {{ .Name }}.tex

{{ .Name }}.pdf: ../../*.tex ../../trademark.pdf *.tex {{ .Name }}.tex
	pdflatex --file-line-error -interaction=nonstopmode {{ .Name }}.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: {{ .Name }}.pdf
	open {{ .Name }}.pdf

o:
	make open

defs: {{ .Name }}.tex
	defs {{ .Name }}.tex

spell:
	aspell -c {{ .Name }}.tex

s:
	make spell

reset:
	rm graph.pdf
	rm graph.csv
	rm graph.graphviz

clean:
	make reset

remake:
	rm {{.Name}}.*
	make reset
	bbk_sheet -mode mk > Makefile`
