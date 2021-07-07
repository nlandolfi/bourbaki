# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make ordered_pair_equality.pdf

ordered_pair_equality.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > ordered_pair_equality.tex

ordered_pair_equality.pdf: ../../*.tex ../../trademark.pdf *.tex ordered_pair_equality.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode ordered_pair_equality.tex
	../../latexrun ordered_pair_equality.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: ordered_pair_equality.pdf
	open ordered_pair_equality.pdf

o:
	make open

terms: ordered_pair_equality.tex
	bbk_sheet -mode ts ordered_pair_equality.tex

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
	rm -f ordered_pair_equality.*
	make reset
	bbk_sheet -mode mk > Makefile