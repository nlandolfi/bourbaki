# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make supremum_norm_complete.pdf
	make graph.pdf

supremum_norm_complete.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > supremum_norm_complete.tex

supremum_norm_complete.pdf: ../../*.tex ../../trademark.pdf *.tex supremum_norm_complete.tex
	# pdflatex --file-line-error -interaction=nonstopmode supremum_norm_complete.tex
	../../latexrun supremum_norm_complete.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: supremum_norm_complete.pdf
	open supremum_norm_complete.pdf

o:
	make open

terms: supremum_norm_complete.tex
	bbk_sheet -mode ts supremum_norm_complete.tex

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
	rm -f supremum_norm_complete.*
	make reset
	bbk_sheet -mode mk > Makefile