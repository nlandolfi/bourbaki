# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make profit_maximizing_production_allocation.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

profit_maximizing_production_allocation.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > profit_maximizing_production_allocation.tex

profit_maximizing_production_allocation.pdf: ../../*.tex ../../trademark.pdf *.tex profit_maximizing_production_allocation.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode profit_maximizing_production_allocation.tex
	../../latexrun profit_maximizing_production_allocation.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: profit_maximizing_production_allocation.pdf
	open profit_maximizing_production_allocation.pdf

o:
	make open

terms: profit_maximizing_production_allocation.tex
	bbk sheet -mode ts profit_maximizing_production_allocation.tex

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
	rm -f profit_maximizing_production_allocation.*
	make reset
	bbk sheet -mode mk > Makefile