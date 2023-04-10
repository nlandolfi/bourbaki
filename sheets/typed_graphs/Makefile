# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make typed_graphs.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

typed_graphs.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > typed_graphs.tex

typed_graphs.pdf: ../../*.tex ../../trademark.pdf *.tex typed_graphs.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode typed_graphs.tex
	../../latexrun typed_graphs.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: typed_graphs.pdf
	open typed_graphs.pdf

o:
	make open

terms: typed_graphs.tex
	bbk sheet -mode ts typed_graphs.tex

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
	rm -f typed_graphs.*
	make reset
	bbk sheet -mode mk > Makefile