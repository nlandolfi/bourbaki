# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make absolute_value.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

absolute_value.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > absolute_value.tex

absolute_value.pdf: ../../*.tex ../../trademark.pdf *.tex absolute_value.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode absolute_value.tex
	../../latexrun absolute_value.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: absolute_value.pdf
	open absolute_value.pdf

o:
	make open

terms: absolute_value.tex
	bbk sheet -mode ts absolute_value.tex

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
	rm -f absolute_value.*
	make reset
	bbk sheet -mode mk > Makefile