# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make generalized_set_dualities.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

generalized_set_dualities.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > generalized_set_dualities.tex

generalized_set_dualities.pdf: ../../*.tex ../../trademark.pdf *.tex generalized_set_dualities.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode generalized_set_dualities.tex
	../../latexrun generalized_set_dualities.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: generalized_set_dualities.pdf
	open generalized_set_dualities.pdf

o:
	make open

terms: generalized_set_dualities.tex
	bbk_sheet -mode ts generalized_set_dualities.tex

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
	rm -f generalized_set_dualities.*
	make reset
	bbk_sheet -mode mk > Makefile