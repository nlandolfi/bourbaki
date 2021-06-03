# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make norm_completeness.pdf

norm_completeness.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > norm_completeness.tex

norm_completeness.pdf: ../../*.tex ../../trademark.pdf *.tex norm_completeness.tex
	# pdflatex --file-line-error -interaction=nonstopmode norm_completeness.tex
	../../latexrun norm_completeness.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: norm_completeness.pdf
	open norm_completeness.pdf

o:
	make open

terms: norm_completeness.tex
	bbk_sheet -mode ts norm_completeness.tex

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
	rm -f norm_completeness.*
	make reset
	bbk_sheet -mode mk > Makefile