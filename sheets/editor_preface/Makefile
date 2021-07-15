# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make editor_preface.pdf

editor_preface.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > editor_preface.tex

editor_preface.pdf: ../../*.tex ../../trademark.pdf *.tex editor_preface.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode editor_preface.tex
	../../latexrun editor_preface.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: editor_preface.pdf
	open editor_preface.pdf

o:
	make open

terms: editor_preface.tex
	bbk_sheet -mode ts editor_preface.tex

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
	rm -f editor_preface.*
	make reset
	bbk_sheet -mode mk > Makefile