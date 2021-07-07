# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make zero_one_law.pdf

zero_one_law.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > zero_one_law.tex

zero_one_law.pdf: ../../*.tex ../../trademark.pdf *.tex zero_one_law.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode zero_one_law.tex
	../../latexrun zero_one_law.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: zero_one_law.pdf
	open zero_one_law.pdf

o:
	make open

terms: zero_one_law.tex
	bbk_sheet -mode ts zero_one_law.tex

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
	rm -f zero_one_law.*
	make reset
	bbk_sheet -mode mk > Makefile