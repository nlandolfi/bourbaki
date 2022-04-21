package bbk

const MakefileTemplate = `# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make {{ .Name }}.pdf
{{ if .HasLitFile }}
sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex
{{ end }}
{{ .Name }}.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > {{ .Name }}.tex

{{ .Name }}.pdf: ../../*.tex ../../trademark.pdf *.tex {{ .Name }}.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode {{ .Name }}.tex
	../../latexrun {{ .Name }}.tex
	make terms

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

terms: {{ .Name }}.tex
	bbk_sheet -mode ts {{ .Name }}.tex

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
	rm -f {{.Name}}.*
	make reset
	bbk_sheet -mode mk > Makefile`
