# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make regular_expressions.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

regular_expressions.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > regular_expressions.tex

regular_expressions.pdf: ../../*.tex ../../trademark.pdf *.tex regular_expressions.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode regular_expressions.tex
	../../latexrun regular_expressions.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: regular_expressions.pdf
	open regular_expressions.pdf

o:
	make open

terms: regular_expressions.tex
	bbk sheet -mode ts regular_expressions.tex

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
	rm -f regular_expressions.*
	make reset
	bbk sheet -mode mk > Makefile