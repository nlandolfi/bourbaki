# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make directed_graphs.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

directed_graphs.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > directed_graphs.tex

directed_graphs.pdf: ../../*.tex ../../trademark.pdf *.tex directed_graphs.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode directed_graphs.tex
	../../latexrun directed_graphs.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: directed_graphs.pdf
	open directed_graphs.pdf

o:
	make open

terms: directed_graphs.tex
	bbk sheet -mode ts directed_graphs.tex

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
	rm -f directed_graphs.*
	make reset
	bbk sheet -mode mk > Makefile