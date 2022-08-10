# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make range_and_null_space.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

range_and_null_space.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > range_and_null_space.tex

range_and_null_space.pdf: ../../*.tex ../../trademark.pdf *.tex range_and_null_space.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode range_and_null_space.tex
	../../latexrun range_and_null_space.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: range_and_null_space.pdf
	open range_and_null_space.pdf

o:
	make open

terms: range_and_null_space.tex
	bbk_sheet -mode ts range_and_null_space.tex

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
	rm -f range_and_null_space.*
	make reset
	bbk_sheet -mode mk > Makefile