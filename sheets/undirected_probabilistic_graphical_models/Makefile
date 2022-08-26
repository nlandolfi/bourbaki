# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make undirected_probabilistic_graphical_models.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

undirected_probabilistic_graphical_models.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > undirected_probabilistic_graphical_models.tex

undirected_probabilistic_graphical_models.pdf: ../../*.tex ../../trademark.pdf *.tex undirected_probabilistic_graphical_models.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode undirected_probabilistic_graphical_models.tex
	../../latexrun undirected_probabilistic_graphical_models.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: undirected_probabilistic_graphical_models.pdf
	open undirected_probabilistic_graphical_models.pdf

o:
	make open

terms: undirected_probabilistic_graphical_models.tex
	bbk_sheet -mode ts undirected_probabilistic_graphical_models.tex

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
	rm -f undirected_probabilistic_graphical_models.*
	make reset
	bbk_sheet -mode mk > Makefile